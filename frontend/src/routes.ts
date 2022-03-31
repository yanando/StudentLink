import Home from "./routes/Home.svelte"
import Login from './routes/Login.svelte'
import Register from './routes/Register.svelte'
import Landing from './routes/Landing.svelte'
import Chat from './routes/Chat.svelte'

export default {
    '/': Home,
    '/login': Login,
    '/register': Register,
    '/landing': Landing,
    '/chat/:id': Chat
}