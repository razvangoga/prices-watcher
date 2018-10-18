# prices-watcher
a small project to watch the prices of some items of interest and learn go in the process

### external libs

* [HTML parsing](https://github.com/antchfx/htmlquery)
* Telemetry - [Azure Application Insights](https://github.com/Microsoft/ApplicationInsights-Go)
* Watched URLs source - file in [Azure Blob Storage](https://github.com/azure/azure-storage-blob-go/)

### deployment
* as Azure Webjob (using *AzureDevOps* to compile / package and deploy)
    * build - manually created build pipeline using the visual designer
    * packaging - Powershell script that creates a zip archive with the expected webjob structure
        * [packaging 1](https://markheath.net/post/managing-webjobs-with-kudu-api)
        * [packaging 2](https://github.com/projectkudu/kudu/wiki/WebJobs)
    * deploy - manually created Azure App Service that deploys the zip created during packaging


