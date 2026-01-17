<script lang="ts">
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import User from "@lucide/svelte/icons/user";
    import Users from "@lucide/svelte/icons/users";
    import FileQuestion from "@lucide/svelte/icons/file-question";
</script>

<div class="flex flex-col items-center justify-center text-center px-4 py-8">
    <div class="neo-border neo-shadow-lg bg-accent px-8 py-4 mb-8 rotate-[1deg]">
        <h1 class="text-3xl md:text-4xl font-black uppercase tracking-tight">
            Working alone or in a group?
        </h1>
    </div>

    <p class="text-lg font-medium max-w-md mb-8">
        This helps us know whose information to collect.
    </p>

    <div class="flex flex-col sm:flex-row gap-6">
        <button
            type="button"
            onclick={() => {
                wizardStore.setTeamType("individual");
                wizardStore.nextStep();
            }}
            class="neo-border neo-shadow bg-card px-10 py-8 text-xl font-black uppercase tracking-wide flex flex-col items-center gap-4 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer hover:bg-secondary"
        >
            <div class="neo-border-sm bg-primary p-4">
                <User class="size-10" />
            </div>
            <span>Just Me</span>
        </button>
        <button
            type="button"
            onclick={() => {
                wizardStore.setTeamType("group");
                wizardStore.nextStep();
            }}
            class="neo-border neo-shadow bg-card px-10 py-8 text-xl font-black uppercase tracking-wide flex flex-col items-center gap-4 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer hover:bg-accent"
        >
            <div class="neo-border-sm bg-secondary p-4">
                <Users class="size-10" />
            </div>
            <span>Group</span>
        </button>
    </div>

    <button
        type="button"
        onclick={() => {
            wizardStore.setTeamType("blank");
            wizardStore.nextStep();
        }}
        class="mt-6 neo-border neo-shadow bg-muted px-6 py-4 flex items-center gap-4 hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all cursor-pointer hover:bg-card"
    >
        <div class="neo-border-sm bg-card p-2">
            <FileQuestion class="size-6" />
        </div>
        <div class="text-left">
            <span class="font-bold uppercase">Leave Blank</span>
            <p class="text-sm text-muted-foreground font-normal normal-case">
                Fill in name & reg number by hand after printing
            </p>
        </div>
    </button>
</div>
