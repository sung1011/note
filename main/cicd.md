# CI CD

## CI

    Continuous Integration 持续集成; 开发完成后自动进行代码检查, 单元测试, 打包部署到测试, 跑集成测试, 跑自动化测试用例; 即负责到开发环境.

- 代码检查
- 单元测试
- 集成测试 自动跑测试用例.

## CD continuous-delivery

    Continuous Delivery 持续交付
    与CD比, 少一步自动部署到线上, 即达到可上线的地步即可. 即负责到上线前.

- 灰度 新业务代码从生产环境替换两个节点

## CD continuous-deploy

    Continuous Deploy 持续部署
    代码测试通过后自动部署到类生产环境, 测试通过后灰度, 最后自动部署到线上. 即负责到上线.

- 发布 将灰度版本逐渐扩大到所有节点
