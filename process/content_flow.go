package process

import (
	"demoProject/internal/dao"
	"encoding/json"
	"fmt"
	v1 "github.com/s8sg/goflow/flow/v1"
	goflow "github.com/s8sg/goflow/v1"
	"gorm.io/gorm"
)

func ExecContentFlow(db *gorm.DB) {
	contentFlow := &contentFlow{
		contentDao: dao.NewContentDao(db),
	}
	fs := goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		WorkerConcurrency: 5,
	}
	_ = fs.Register("content-flow", contentFlow.flowHandle)
	_ = fs.Start()
}

type contentFlow struct {
	contentDao *dao.ContentDao
}

func (c *contentFlow) flowHandle(flow *v1.Workflow, context *v1.Context) error {
	dag := flow.Dag()
	dag.Node("input", c.input)
	dag.Node("verify", c.verify)
	dag.Node("category", c.category)
	branches := dag.ConditionalBranch("branchs",
		[]string{"category", "branch", "thumbnail", "format", "fail"},
		func(bytes []byte) []string {
			var data map[string]interface{}
			if err := json.Unmarshal(bytes, &data); err != nil {
				return nil
			}
			if data["approval_status"].(float64) == 2 {
				return []string{"category", "branch", "thumbnail", "format", "fail"}
			}
			return []string{"fail"}
		}, v1.Aggregator(func(m map[string][]byte) ([]byte, error) {
			return []byte("ok"), nil
		}))

	branches["category"].Node("branch", c.category)
	branches["thumbnail"].Node("thumbnail", c.thumbnail)
	branches["pass"].Node("category", c.pass)
	branches["format"].Node("format", c.format)
	branches["fail"].Node("fail", c.fail)

	dag.Edge("input", "verify")
	dag.Edge("verify", "branches")
	dag.Edge("branches", "finish")
	return nil
}

func (c *contentFlow) input(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec input")
	var input map[string]int
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	id := input["content_id"]
	detail, err := c.contentDao.First(id)
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(map[string]interface{}{
		"title":      detail.Title,
		"video_url":  detail.VideoURL,
		"content_id": detail.ID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *contentFlow) verify(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec verify")
	var detail map[string]interface{}
	if err := json.Unmarshal(data, &detail); err != nil {
		return nil, err
	}
	var (
		title    = detail["title"]
		videoUrl = detail["video_url"]
		id       = detail["content_id"]
	)

	if int(id.(float64))%2 == 0 {
		detail["approval_status"] = 3
	} else {
		detail["approval_status"] = 2
	}
	fmt.Println(id, title, videoUrl)
	return json.Marshal(detail)
}

func (c *contentFlow) category(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec category")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int(input["content_id"].(float64))
	if err := c.contentDao.UpdateById(contentId, "category", "category"); err != nil {
		return nil, err
	}
	return []byte("category"), nil
}

func (c *contentFlow) thumbnail(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec thumbnail")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int(input["content_id"].(float64))
	if err := c.contentDao.UpdateById(contentId, "thumbnail", "thumbnail"); err != nil {
		return nil, err
	}
	return []byte("format"), nil
}

func (c *contentFlow) format(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec format")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int(input["content_id"].(float64))
	if err := c.contentDao.UpdateById(contentId, "approval_status", 2); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *contentFlow) pass(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec pass")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int(input["content_id"].(float64))
	if err := c.contentDao.UpdateById(contentId, "pass", "pass"); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *contentFlow) fail(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Println("exec pass")
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}
	contentId := int(input["content_id"].(float64))
	if err := c.contentDao.UpdateById(contentId, "pass", "pass"); err != nil {
		return nil, err
	}
	return data, nil
}
