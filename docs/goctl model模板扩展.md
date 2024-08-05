# goctl model模板扩展
## 原因
在`goctl`默认的模板实现中，只会生成下面的几个接口
```go
transactionModel interface {
    Insert(ctx context.Context, data *Transaction) (sql.Result, error)
    FindOne(ctx context.Context, id int64) (*Transaction, error)
    Update(ctx context.Context, data *Transaction) error
    Delete(ctx context.Context, id int64) error
}
```
但是，我们的业务需要更多的接口，比如说`FindAll`: 查询所有记录。上面的四个接口是不够看的，需要对模板进行扩展。

## 步骤
1. 首先确定我们使用的`goctl版本`。
```sh
goctl --version
```
> 执行以上命令，在我的开发机器上，会出现`goctl version 1.7.0 windows/amd64`的结果。
2. 下载这个版本的模板到本地，执行下面命令即可
```sh
goctl template init --home=./goctl/1.7.0
```
> 命令执行成功后，会在当前执行命令的目录下生成`goctl/1.7.0`目录，然后在`1.7.0`目录中就包含了各个模板文件
3. 查看默认的模板文件
```sh
PS D:\YINC_DEVELOPMENT\go\stock\goctl\1.7.0> tree /F
文件夹 PATH 列表
卷序列号为 0001-C64B     
D:.
├─api
│      config.tpl        
│      context.tpl       
│      etc.tpl
│      handler.tpl       
│      logic.tpl
│      main.tpl
│      middleware.tpl    
│      route-addition.tpl
│      routes.tpl        
│      template.tpl      
│      types.tpl
│      
├─docker
│      docker.tpl
│
├─gateway
│      etc.tpl
│      main.tpl
│
├─kube
│      deployment.tpl
│      job.tpl
│
├─model
│      customized.tpl
│      delete.tpl
│      err.tpl
│      field.tpl
│      find-one-by-field-extra-method.tpl
│      find-one-by-field.tpl
│      find-one.tpl
│      import-no-cache.tpl
│      import.tpl
│      insert.tpl
│      interface-delete.tpl
│      interface-find-one-by-field.tpl
│      interface-find-one.tpl
│      interface-insert.tpl
│      interface-update.tpl
│      model-gen.tpl
│      model-new.tpl
│      model.tpl
│      table-name.tpl
│      tag.tpl
│      types.tpl
│      update.tpl
│      var.tpl
│
├─mongo
│      err.tpl
│      model.tpl
│      model_custom.tpl
│      model_types.tpl
│
├─newapi
│      newtemplate.tpl
│
└─rpc
        call.tpl
        config.tpl
        etc.tpl
        logic-func.tpl
        logic.tpl
        main.tpl
        server-func.tpl
        server.tpl
        svc.tpl
        template.tpl
```