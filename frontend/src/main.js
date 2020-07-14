import 'core-js/stable';
import 'regenerator-runtime/runtime';
import "materialize-css/dist/css/materialize.min.css";
import 'materialize-css';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import VueDataTables from 'vue-data-tables';
import lang from 'element-ui/lib/locale/lang/en';
import locale from 'element-ui/lib/locale';
import Vuelidate from 'vuelidate';
import Vue from 'vue';
import App from './App.vue';
import PerfectScrollbar from 'vue2-perfect-scrollbar'
import 'vue2-perfect-scrollbar/dist/vue2-perfect-scrollbar.css'


Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Vue.use(Vuelidate)
Vue.use(ElementUI)
Vue.use(VueDataTables)
locale.use(lang)
Vue.use(PerfectScrollbar)

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
