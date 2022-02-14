import { createApp } from 'vue';
import App from './App';
import router from './router';
import store from './store';
import 'virtual:windi.css';

const app = createApp(App);

app.use(router);
app.use(store);

const meta = document.createElement('meta');
meta.name = 'naive-ui-style';
document.head.appendChild(meta);

app.mount('#app');
