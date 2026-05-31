FROM node:20

#
WORKDIR /app

#依存パッケージのインストール
COPY package.json package-lock.json .
RUN npm install
COPY . .

# viteのデフォルトポート
EXPOSE 5173

# コンテナ起動時のコマンド
CMD ["npm", "run", "dev"]