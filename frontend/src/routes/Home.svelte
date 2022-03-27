<script lang="ts">
  import { push } from "svelte-spa-router";
  setTimeout(() => {
    push("/login");
  }, 1000);
</script>

<div id="wrapper">
  <div class="lds-dual-ring" />
</div>

<style lang="scss">
  #wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100vw;
    height: 100vh;
    background-color: #252427;
  }

  .lds-dual-ring {
    display: inline-block;
    width: 80px;
    height: 80px;
  }
  .lds-dual-ring:after {
    content: " ";
    display: block;
    width: 64px;
    height: 64px;
    margin: 8px;
    border-radius: 50%;
    border: 6px solid #0c7d9d;
    border-color: #0c7d9d transparent #0c7d9d transparent;
    animation: lds-dual-ring 1.2s linear infinite;
  }
  @keyframes lds-dual-ring {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
