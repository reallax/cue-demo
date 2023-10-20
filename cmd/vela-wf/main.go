package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"cuelang.org/go/cue/format"
	"cuelang.org/go/cue"
	// "github.com/kubevela/workflow/api/v1alpha1"
	"github.com/kubevela/workflow/pkg/cue/model/value"
	// "github.com/kubevela/workflow/pkg/tasks/custom"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/kubevela/workflow/pkg/cue/model/sets"
)

func main() {
	props := &runtime.RawExtension{Raw: []byte(`
{"url": "http://www.baidu.com", "method": "GET"}
		`)}
	
	pstr, err := getParameterTemplate(props)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pstr)

	// 模拟 pkg/tasks/custom/task.go#MakeBasicValue()
	paramStr := "parameter: {}\n"
	if pstr != "" {
		paramStr = fmt.Sprintf("parameter: {%s}\n", pstr)
	}

	ctx := map[string]interface{}{
		"name":     "test-xxxxxxxxx",
		"appName":  "symphony-aabb-xx",
		"revision": "revision.hash.cs1fHqResQ",
	}
	template := strings.Join([]string{getContextTemplate(ctx), paramStr}, "\n")

	vars, err := value.NewValue("", nil, "")
	if err != nil {
		return
	}
	// 模拟 pkg/tasks/custom/task.go#MakeParameter()
	if template == "" {
		template = "{}"
	}
	basicVal, err := vars.MakeValue(template)
	if err != nil {
		return
	}

	basicTemplate, err := basicVal.String()
	if err != nil {
		return
	}
	taskv, err := value.NewValue(strings.Join([]string{stepRequestTmpl, basicTemplate}, "\n"), nil, "", value.ProcessScript, value.TagFieldOrder)
	if err != nil {
		return
	}

	// 三种输出value的方式。
	// 1 没渲染出#do属性
	// 2 和 3 效果相同，且均能渲染出#do属性
	fmt.Printf("##1. fmt taskv##:\n%# v\n", taskv.CueValue())

	ss, _ := sets.ToString(taskv.CueValue())
	fmt.Printf("##2.fmt taskv##:\n%s\n", ss)

	syn := taskv.CueValue().Syntax(
		cue.Final(),         // close structs and lists
		cue.Concrete(false), // allow incomplete values
		cue.Definitions(false),
		cue.Hidden(true),
		cue.Optional(true),
		cue.Attributes(true),
		cue.Docs(true),
		// cue.All(),
	)

	bs, _ := format.Node(
		syn,
		// format.TabIndent(false),
		// format.UseSpaces(2),
	)

	// print to stdout
	fmt.Println("##3. fmt taskv##: \n" + string(bs))

	do := opTpy(taskv)
	fmt.Printf("  #do: %s\n", do)

	provider := opProvider(taskv)
	fmt.Printf("  #provider: %s\n", provider)

	// iter, _ := taskv.CueValue().List();
	// cueIter, err := iter.target.v.Fields(cue.Definitions(true), cue.Hidden(true), cue.All())
	// for iter.Next() {
	// 	l := iter.Label()
	// 	v := iter.Value()

	// 	fmt.Printf(" label: %s  v: %# v\n", l, v)
	// }
}

// 模拟 getParameterTemplate()
func getParameterTemplate(props *runtime.RawExtension) (string, error) {
	if props != nil && len(props.Raw) > 0 {
		params := map[string]interface{}{}
		bt, err := props.MarshalJSON()
		if err != nil {
			return "", err
		}
		if err := json.Unmarshal(bt, &params); err != nil {
			return "", err
		}
		b, err := json.Marshal(params)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", nil
}

func getContextTemplate(ctx map[string]interface{}) string {
	d, err := json.Marshal(ctx)
	if err != nil {
		return ""
	}
	buff := fmt.Sprintf("\n %s", structMarshal(string(d)))
	c := fmt.Sprintf("context: %s", structMarshal(buff))

	contextTempl := ""
	contextTempl += "\n" + c
	return contextTempl
}

func structMarshal(v string) string {
	skip := false
	v = strings.TrimFunc(v, func(r rune) bool {
		if !skip {
			if unicode.IsSpace(r) {
				return true
			}
			skip = true

		}
		return false
	})

	if strings.HasPrefix(v, "{") {
		return v
	}
	return fmt.Sprintf("{%s}", v)
}

func opTpy(v *value.Value) string {
	return getLabel(v, "http.#do")
}

func opProvider(v *value.Value) string {
	provider := getLabel(v, "#provider")
	if provider == "" {
		provider = "builtin"
	}
	return provider
}

func getLabel(v *value.Value, label string) string {
	do, err := v.Field(label)
	if err == nil && do.Exists() {
		if str, err := do.String(); err == nil {
			return str
		}
	}
	return ""
}

const stepRequestTmpl = `import (
	"vela/op"
	"encoding/json"
)

http: op.#HTTPDo & {
	method: parameter.method
	url:    parameter.url
	request: {
		if parameter.body != _|_ {
			body: json.Marshal(parameter.body)
		}
		if parameter.header != _|_ {
			header: parameter.header
		}
	}
}
fail: op.#Steps & {
	if http.response.statusCode > 400 {
		requestFail: op.#Fail & {
			message: "request of \(parameter.url) is fail: \(http.response.statusCode)"
		}
	}
}
response: json.Unmarshal(http.response.body)
parameter: {
	url:    string
	method: *"GET" | "POST" | "PUT" | "DELETE"
	body?: {...}
	header?: [string]: string
}`
