# Программа для конвертации PDF в изображения

## Как это работает:
- Программа принимает входной PDF-файл и (опционально) выходную директорию.
- Использует go-fitz для рендеринга каждой страницы PDF в изображение.
- Сохраняет каждую страницу в формате PNG в папке с названием исходного файла.
