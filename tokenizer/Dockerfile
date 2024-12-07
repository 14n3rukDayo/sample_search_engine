# for development
FROM python:3.9
WORKDIR /app

COPY . .
RUN pip install --no-cache-dir --upgrade pip setuptools
RUN pip install "fastapi[standard]"
RUN pip install "pydantic"
RUN pip install -U "ginza" "ja_ginza"





EXPOSE 5000

# アプリケーションを実行
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "5000", "--reload"]