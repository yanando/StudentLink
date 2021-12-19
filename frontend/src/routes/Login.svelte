<script lang="ts">
    let username: string
    let password: string

    const submit = async () => {
        if (!username || !password) {
            alert('Geen gebruikersnaam of wachtwoord ingevoerd!')
            return
        }

        const payload = {
            username: username,
            password: password
        }

        const resp = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: {
                'content-type': 'application/json'
            }
        })

        console.log(resp)
    }
</script>

<div id="wrapper">
    <div class="inside-wrapper">
        <h1>StudentLink Login</h1>
        <div class="input-wrapper">
            <input class="input" type="text" name="username" placeholder="Username" bind:value={username}>
            <input class="input" type="password" name="password" placeholder="Password" bind:value={password}>
        </div>
        <button id="submit" on:click={submit}>Login</button>
    </div>
</div>

<style lang="scss">
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