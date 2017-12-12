<template>
  <div class="siimple-grid">
    <div class="siimple-grid-row wrapper">
      <div class="siimple-grid-col siimple-grid-col--8 siimple-grid-col-md--12">
          <Converter :rate="rate" />
          <Information :rate="rate" />
          <Graph :rate="rate" />
      </div>
    </div>
  </div>
</template>

<script>
import Converter from './Converter';
import Information from './Information';
import Graph from './Graph';

export default {
  name: 'ZnyTo',
  components: {
    Converter,
    Information,
    Graph,
  },
  data() {
    return {
      rate: 0,
    };
  },
  mounted() {
    this.getRate();
  },
  beforeDestroy() {
    clearInterval(this.interval);
  },
  methods: {
    // TODO: replace with WebSockets
    getRate() {
      this.interval = setInterval(() => {
        const url = 'https://46da7e28.ngrok.io/';
        const xhr = new XMLHttpRequest();
        xhr.open('GET', url);
        xhr.addEventListener('load', () => {
          this.rate = parseFloat(xhr.responseText);
        }, false);
        xhr.send();
      }, 1000);
    },
  },
};
</script>

<style scoped>
.wrapper {
  display: flex;
  justify-content: center;
}
</style>
