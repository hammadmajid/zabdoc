<script lang="ts">
    import SunIcon from "@lucide/svelte/icons/sun";
    import MoonIcon from "@lucide/svelte/icons/moon";
    import MonitorIcon from "@lucide/svelte/icons/monitor";
    import { setMode } from "mode-watcher";
    import { userPrefersMode } from "mode-watcher";
    import * as Select from "$lib/components/ui/select/index.js";
    import Label from "$lib/components/ui/label/label.svelte";

    const modeOptions = [
        { value: "light", label: "Light", icon: SunIcon },
        { value: "dark", label: "Dark", icon: MoonIcon },
        { value: "system", label: "System", icon: MonitorIcon },
    ];

    let selectedMode = $state("");

    // Initialize selectedMode when component mounts
    $effect(() => {
        selectedMode = userPrefersMode.current;
    });

    function handleModeChange(value: string | undefined) {
        if (
            value &&
            (value === "light" || value === "dark" || value === "system")
        ) {
            setMode(value);
        }
    }
</script>

<div class="space-y-2">
    <Label for="theme-select">Theme</Label>
    <Select.Root
        type="single"
        name="theme"
        value={selectedMode}
        onValueChange={handleModeChange}
    >
        <Select.Trigger id="theme-select" class="w-full">
            <div class="flex items-center gap-2">
                {#each modeOptions as option}
                    {#if option.value === selectedMode}
                        <option.icon />
                        <span>{option.label}</span>
                    {/if}
                {/each}
            </div>
        </Select.Trigger>
        <Select.Content>
            <Select.Group>
                {#each modeOptions as option (option.value)}
                    <Select.Item value={option.value} label={option.label}>
                        <div class="flex items-center gap-2">
                            <option.icon />
                            <span>{option.label}</span>
                        </div>
                    </Select.Item>
                {/each}
            </Select.Group>
        </Select.Content>
    </Select.Root>
</div>
