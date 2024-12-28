# ベースイメージを定義
FROM node:20-alpine AS base

WORKDIR /app

# package.jsonとlockファイルをコピーして依存関係をインストール（キャッシュ活用）
COPY package.json package-lock.json ./
RUN npm install --only=production

# Development Stage
FROM base AS development

# 開発用に全依存関係をインストール
RUN npm install

COPY . ./

# 開発用のコマンド
CMD ["npm", "run", "dev"]

# Production Stage
FROM base AS production

COPY . ./

# ビルドを実行（本番用）
RUN npm run build

# 本番用設定を適用
COPY .env.production .env

# 必要なポートを公開
EXPOSE 3000

# 本番用のコマンド
CMD ["npm", "start"]
