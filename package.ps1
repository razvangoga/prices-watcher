$jobName = "prices-watcher"

Add-Type -assembly  System.IO.Compression.FileSystem
$archive = [System.IO.Compression.ZipFile]::Open("$pwd/$jobName.zip","Create")

$exe = "$pwd/$jobName.exe"
$settings = "$pwd/settings.json"

$jobPath = "\app_data\jobs\triggered\$jobName"

$zexe = [System.IO.Compression.ZipFileExtensions]::CreateEntryFromFile($archive,$exe,"$jobPath\prices-watcher.exe","optimal")
$zsettings = [System.IO.Compression.ZipFileExtensions]::CreateEntryFromFile($archive,$settings,"$jobPath\settings.json","optimal")

$archive.Dispose()

Write-Host "Done..."