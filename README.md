网络测试工具

wails初始化
wails init -n myproject -t vue

开发
wails dev
http://localhost:34115/
http://localhost:5173/

构建
cd frontend
rm -rf node_modules
rm package-lock.json
cnpm install
wails build

安装vue
npm install element-plus --save
npm install @element-plus/icons-vue

