<script lang="ts">
    import '../app.css';
    import favicon from '$lib/assets/favicon.svg';

    let {children} = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>

<div class="max-w-7xl mx-auto h-screen">
    {@render children?.()}
</div>
