Add-Type -assembly  System.IO.Compression.FileSystem
$zip =  [System.IO.Compression.ZipFile]::Open("prices-wacher.zip","Create")