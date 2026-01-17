<script lang="ts">
    import Calendar from "$lib/components/ui/calendar/calendar.svelte";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import ChevronDownIcon from "@lucide/svelte/icons/chevron-down";
    import CalendarIcon from "@lucide/svelte/icons/calendar";
    import { formStore } from "$lib/stores/form-store.svelte";
    import {
        getLocalTimeZone,
        today,
        type CalendarDate,
    } from "@internationalized/date";

    let open = $state(false);
    let value = $state<CalendarDate | undefined>(today(getLocalTimeZone()));

    const formattedDate = $derived(
        value
            ? value.toDate(getLocalTimeZone()).toLocaleDateString("en-US", {
                  year: "numeric",
                  month: "short",
                  day: "numeric",
              })
            : ""
    );

    const displayDate = $derived(
        value
            ? value.toDate(getLocalTimeZone()).toLocaleDateString("en-US", {
                  year: "numeric",
                  month: "short",
                  day: "numeric",
              })
            : "Select due date"
    );

    // Sync with store when date changes
    $effect(() => {
        formStore.setDocumentDate(formattedDate);
    });
</script>

<div class="flex flex-col gap-3">
    <Popover.Root bind:open>
        <Popover.Trigger id="date">
            {#snippet child({ props })}
                <Button
                    {...props}
                    variant="outline"
                    class="w-full justify-between font-normal {!value
                        ? 'text-muted-foreground'
                        : ''}"
                >
                    <div class="flex items-center gap-2">
                        <CalendarIcon class="h-4 w-4" />
                        {displayDate}
                    </div>
                    <ChevronDownIcon class="h-4 w-4 opacity-50" />
                </Button>
            {/snippet}
        </Popover.Trigger>
        <Popover.Content class="w-auto overflow-hidden p-0" align="start">
            <Calendar
                type="single"
                bind:value
                captionLayout="dropdown"
                onValueChange={() => {
                    open = false;
                }}
            />
        </Popover.Content>
    </Popover.Root>
</div>
