<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button";
</script>

<div class="flex flex-col items-center justify-center text-center h-full px-4">
    <h1 class="text-5xl font-bold tracking-tight">zabdoc</h1>
    <p class="mt-4 text-lg text-muted-foreground max-w-xl">
        Generate professional assignment cover sheets and lab tasks for SZABIST
        students
    </p>

    <div class="mt-8 flex gap-4">
        <a href="/assignment" class={buttonVariants({ variant: "default" })}>
            Assignment Cover
        </a>
        <Button disabled>Lab Task</Button>
    </div>
</div>
