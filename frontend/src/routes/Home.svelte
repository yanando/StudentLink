<script lang="ts">
    import {link} from 'svelte-spa-router'
</script>

<h1 id="title">StudentLink</h1>
<div id="wrapper">
    <a href="/login" id="login-button" use:link>LOGIN</a>
</div>

<style lang="scss">
    body {
        overflow: hidden;
    }

    #title {
        text-align: center;
    }

    #wrapper {
        display: flex;
        justify-content: center;
    }
</style>