import Home from "./routes/Home.svelte"
import Login from './routes/Login.svelte'
import Profile from './routes/Profile.svelte'

export default {
    '/': Home,
    '/login': Login,
    '/profile': Profile
}