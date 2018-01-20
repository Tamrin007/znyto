<template>
  <highcharts :options="options"></highcharts>
</template>

<script>
function utc2dateString(msec) {
  const d = new Date();
  d.setTime(msec);
  return `${d.getHours()}:${d.getMinutes()}:${d.getSeconds()}`;
}

export default {
  name: 'Information',
  props: {
    rate: Number,
  },
  data() {
    return {
      options: {
        rangeSelector: {
          selected: 1,
        },
        title: {
          text: 'ZNY to JPY',
        },
        series: [{
          name: 'JPY/ZNY',
          data: [],
          tooltip: {
            valueDecimals: 2,
          },
          animation: false,
        }],
        xAxis: [{  // X軸
          labels: {
            formatter() {
              return utc2dateString(this.value);
            },
          },
        }],
      },
    };
  },
  watch: {
    rate(r) {
      const now = Date.now();
      /* eslint-disable no-console */
      this.options.series[0].data.push([now, r]);
      console.log(this.options.series[0].data);
    },
  },
};
</script>
