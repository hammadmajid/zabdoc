<script lang="ts">
    import '../app.css';
    import favicon from '$lib/assets/favicon.svg';

    import {ModeWatcher} from "mode-watcher";

    let {children} = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
    <title>zabcover</title>
</svelte:head>

<ModeWatcher/>
{@render children?.()}
