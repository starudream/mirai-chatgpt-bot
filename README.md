# Mirai ChatGPT Bot

![Golang](https://img.shields.io/github/actions/workflow/status/starudream/mirai-chatgpt-bot/golang.yml?label=golang&style=for-the-badge)
![Docker](https://img.shields.io/github/actions/workflow/status/starudream/mirai-chatgpt-bot/docker.yml?label=docker&style=for-the-badge)
![Release](https://img.shields.io/github/v/release/starudream/mirai-chatgpt-bot?include_prereleases&sort=semver&style=for-the-badge)
![License](https://img.shields.io/github/license/starudream/mirai-chatgpt-bot?style=for-the-badge)

## Config

```yaml
debug: true
log:
  level: debug
addr: :8080
openai:
  api_key: abc
mirai:
  bot_qq: 000
  target_groups:
    - 222
    - 333
```

## Docker

![Version](https://img.shields.io/docker/v/starudream/mirai-chatgpt-bot?sort=semver&style=for-the-badge)
![Size](https://img.shields.io/docker/image-size/starudream/mirai-chatgpt-bot?sort=semver&style=for-the-badge)
![Pull](https://img.shields.io/docker/pulls/starudream/mirai-chatgpt-bot?style=for-the-badge)

```bash
docker pull starudream/mirai-chatgpt-bot
```

```bash
docker run -d \
    --name mcb \
    --restart always \
    -v /opt/docker/mcb/config.yaml:/config.yaml \
    starudream/mirai-chatgpt-bot
```

## License

[Apache License 2.0](./LICENSE)
