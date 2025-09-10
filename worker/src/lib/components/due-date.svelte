<script lang="ts">
    import Calendar from "$lib/components/ui/calendar/calendar.svelte";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import ChevronDownIcon from "@lucide/svelte/icons/chevron-down";
    import {
        getLocalTimeZone,
        type CalendarDate,
    } from "@internationalized/date";

    let open = $state(false);
    let value = $state<CalendarDate | undefined>();

    const formattedDate = $derived(
        value
            ? value.toDate(getLocalTimeZone()).toISOString().split("T")[0]
            : "",
    );
</script>

<div class="flex flex-col gap-3">
    <input type="hidden" name="date" value={formattedDate} />
    <Popover.Root bind:open>
        <Popover.Trigger id="date">
            {#snippet child({ props })}
                <Button
                    {...props}
                    variant="outline"
                    class="w-full justify-between font-normal"
                >
                    {value
                        ? value.toDate(getLocalTimeZone()).toLocaleDateString()
                        : "Due date"}
                    <ChevronDownIcon />
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
