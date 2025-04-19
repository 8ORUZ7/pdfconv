a simple cli tool written in go that converts pdf files to either `.txt` or `.docx` format.

#### 1. clone the repo or download the script:
```bash
git clone https://github.com/8ORUZ7/pdfconv.git
cd pdf-converter
```

#### 2. install dependencies:
```bash
go mod tidy
```

#### 3. build the executable:
```bash
go build -o build/pdf-converter
```

#### 4. run the app:
```bash
./build/pdf-converter
```

- the path to your pdf file (e.g., `example.pdf`)
- the output format (`txt` or `docx`)

output files will be saved in the `output/` folder.


### 📦 dependencies

- [unidoc/unipdf](https://github.com/unidoc/unipdf) – pdf parsing
- [baliance/gooxml](https://github.com/baliance/gooxml) – docx creation

