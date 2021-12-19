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
    <h1>StudentLink Login</h1>
    <div>
        <input type="text" name="username" placeholder="Username" bind:value={username}>
        <input type="password" name="password" placeholder="Password" bind:value={password}>
    </div>
    <button id="submit" on:click={submit}>Login</button>
</div>

<style>
    #wrapper {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        height: 100vh;
    }
    
    #submit {
        margin-top: 5px;
    }
</style>