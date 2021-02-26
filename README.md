cloudevents-display
===

介绍
---

CloudEvents 规范的 HTTP 事件消费者。

使用 google ko 构建本地镜像
---

```
# Download https://github.com/google/ko/releases/tag/v0.8.0

git clone https://github.com/serverless-x/cloudevents-display.git
cd cloudevents-display

KO_DOCKER_REPO=ko.local ko publish --preserve-import-paths --tags latest ./
```

使用 google ko 发布镜像
---

```
KO_DOCKER_REPO=registry-jinan-lab.inspurcloud.cn/library \
  ko publish --preserve-import-paths --tags latest ./
```
