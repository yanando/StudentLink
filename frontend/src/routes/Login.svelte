<script lang="ts">
    import { user, userStore } from "../userSession";
    import {push} from 'svelte-spa-router'
    import { API_BASE } from "../config";

    let username: string
    let password: string

    let errorMessage: string

    const submit = async () => {
        if (!username || !password) {
            errorMessage = 'Geen gebruikersnaam of wachtwoord ingevoerd!'
            return
        }

        const payload = {
            username: username,
            password: password
        }

        const resp = await fetch(API_BASE+'/api/login', {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: {
                'content-type': 'application/json',
            }
        })

        if (resp.status != 200) {
            errorMessage = (await resp.json()).message
            return
        }

        const userJSON: user = await resp.json()
        userJSON.authToken = resp.headers.get('Auth')

        userStore.set(userJSON)
        
        push('/landing')
    }
</script>

<div id="wrapper">
    <div class="inside-wrapper">
        <h1>StudentLink Login</h1>
        <div class="input-wrapper">
            <p class="error" class:hidden={!errorMessage}>{errorMessage}</p>
            <input class="input" type="text" name="username" placeholder="Username" bind:value={username}>
            <input class="input" type="password" name="password" placeholder="Password" bind:value={password}>
        </div>
        <button id="submit" on:click={submit}>Login</button>
    </div>
</div>

<style lang="scss">
    body {
        overflow: hidden;
    }

    #wrapper {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        
        font-size: 1.5vmin;
        background-color: #252427;
    }

    .inside-wrapper {
        display: flex;
        flex-direction: column;
        justify-content: space-evenly;
        align-items: center;
        height: 50vh;
        width: 25vw;

        color: #7dafb9;
        background-color: #1f1f1e;
    }

    .input-wrapper {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        height: 10vh;
        width: 15vw;
    }

    .hidden {
        display: none;
    }

    .error {
        color: red;
    }

    .input {
        width: 100%;
        height: 40%;
        text-align: center;
        outline: none;
        border: none;

        &::placeholder {
            color: white;
        }

        color: white;
        font-size: 2vmin;
        border-radius: 5px;
        background-color: #252427;
    }
    
    #submit {
        width: 15vw;
        height: 15%;
        outline: none;

        font-size: 2vmin;
        border: none;
        border-radius: 5px;
        color: white;
        background-color: #0c7d9d;

        &:hover {
            cursor: pointer;
            transition: 0.3s ease all;
            background-color: rgba(12,125,157,0.75);
        }
    }
</style>