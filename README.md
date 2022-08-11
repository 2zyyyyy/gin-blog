# 个人博客项目
`golang` `gin` `vue` `javascript`

### 踩坑记录
记录一下项目过程中遇到的问题以及对应的处理方案
1. 创建用户接口一直返回空指针（接口状态500）
  
   - 问题
     ![image-20220802165301243](https://tva1.sinaimg.cn/large/e6c9d24ely1h4sitodbmnj21l30u0gws.jpg)
   
   - 原因
   
     数据库字段类型有调整，虽然mysql.db中设置了自动迁移，但是db的init放在main.go的路由初始化下面，导致没有实际执行
   
     ```go
     func main() {
       routers.InitRouter()
     	// mysql 初始化
     	model.InitDB()
     }
     
     ```
   
   - 解决
     调整main中初始化的顺序（先db后router）
   
2. gorm的preload查询区分大小写

   - 问题：现有结构体`Category`,通过gorm的preload查询数据

     ```sql
     // 用主键检索
     err := db.Preload("category").Table("article").First(&article, id).Error
     ```

   - 查询结果
     ![image-20220811154948437](https://tva1.sinaimg.cn/large/e6c9d24ely1h52vkpebtwj20u013qq64.jpg)

   - 解决
     调整查询语句中preload的结构体名称为大写。

     ```sql
     // 用主键检索
     	err := db.Preload("Category").Table("article").First(&article, id).Error
     ```

     ![image-20220811155235292](https://tva1.sinaimg.cn/large/e6c9d24ely1h52vnjz7vfj20u00widiy.jpg)

3. 1

4. 1
