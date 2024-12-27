/// <reference path="./.sst/platform/config.d.ts" />

const domain = "judethings.com"

export default $config({
  app(input) {
    return {
      name: "sst-go-template",
      removal: "remove",
      home: "aws",
    };
  },
  async run() {
    const { stage } = $app
    const environment = {
      STAGE: stage,
      LOG_LEVEL: process.env.LOG_LEVEL
    }

    const api = new sst.aws.Function("go-api", {
      runtime: "go",
      architecture: "arm64",
      memory: '500 MB',
      timeout: '10 minutes',
      url: { cors: true },
      handler: "./cmd/lambda/api/main.go",
      logging: {
        retention: '1 week',
        format: 'json'
      },
      environment
    })

    const router = new sst.aws.Router("go-router", {
      invalidation: false,
      routes: { "/*": api.url },
      domain: {
        name: `api.go-example.${domain}`,
        redirects: [`www.api.go-example.${domain}`]
      }
    })

    return {
      api: api.url,
      domain: router.url
    }
  }
});
