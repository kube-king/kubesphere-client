# kubesphere-client
kubesphere client package


```go

    c, _ := client.NewDefaultClient(client.WithHost(""), client.WithPasswordAuth("", ""))
    
        d, err := c.Deployment().Get("your cluster", "your namespace", "your deployment")
        if err != nil {
            return
        }
        //fmt.Println(d.Spec.Template.Labels)
        d.Spec.Template.Labels["deploy"] = fmt.Sprintf("%v", time.Now().UnixMilli())
    
        err = c.Deployment().Update("your cluster", "your namespace", "your deployment", d)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
```