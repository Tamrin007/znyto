// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import VueHighcharts from 'vue-highcharts';
import Highcharts from 'highcharts';
import loadStock from 'highcharts/modules/stock';
import loadHighchartsMore from 'highcharts/highcharts-more';
import App from './App';

Vue.config.productionTip = false;
loadStock(Highcharts);
loadHighchartsMore(Highcharts);
Vue.use(VueHighcharts, { Highcharts });

/* eslint-disable no-new */
new Vue({
  el: '#app',
  template: '<App/>',
  components: { App },
});
