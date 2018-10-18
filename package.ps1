$jobName = "prices-watcher"

Add-Type -assembly  System.IO.Compression.FileSystem
$archive = [System.IO.Compression.ZipFile]::Open("$pwd/$jobName.zip","Create")

$exe = "$pwd/$jobName.exe"
$exeZipPath = "\app_data\jobs\triggered\$jobName"

$settings = "$pwd/settings.json"

$zexe = [System.IO.Compression.ZipFileExtensions]::CreateEntryFromFile($archive,$exe,"$exeZipPath\prices-watcher.exe","optimal")
$zsettings = [System.IO.Compression.ZipFileExtensions]::CreateEntryFromFile($archive,$settings,"settings.json","optimal")

$archive.Dispose()

Write-Host "Done..."