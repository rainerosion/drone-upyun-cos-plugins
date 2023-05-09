# drone plugins for upyun cos

## 参数说明
| 参数             | 说明     | 备注          |
| ---------------- |--------|-------------|
| up_operator      | 操作员名称  |             |
| up_password      | 操作员密码  |             |
| up_bucket        | bucket | 服务名称        |
| local_base_path  | 本地路径   | 文件或文件夹      |
| remote_base_path | 对象存储路径 | 文件存放路径（文件夹） |

## 使用方式
````yaml
  - name: upload-upyun
    image: rainteam/upcos:latest
    settings:
      up_operator:
        from_secret: up_operator
      up_password:
        from_secret: up_password
      up_bucket:
        from_bucket: up_bucket
      local_base_path: wechatbot
      remote_base_path: gitea/devops
````