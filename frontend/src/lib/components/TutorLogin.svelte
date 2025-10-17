<script lang="ts">
    import { API_BASE_URL } from "../consts";
    import { userStationStore, userStore } from "../stores";

    let active = $state(false);
    // svelte-ignore non_reactive_update
    let inputRef: HTMLInputElement | undefined;

    async function checkToken(token: string) {
        if(inputRef) {
            let response = await fetch(API_BASE_URL + "/api/token", {
                method: "POST",
                body: token,
            }).then(res => res.json());
            if(response.status) {
                userStore.set(token);
                userStationStore.set(response.data)
                localStorage.setItem("token", token);
                localStorage.setItem("station", response.data)
            } else {
                inputRef.value = "Falsch";
            }
            inputRef.disabled = false;
        }
    }

    function loginClick(e: Event) {
        e.preventDefault();
        if(inputRef) {
            inputRef.disabled = true;
            let token = inputRef.value;
            checkToken(token);
        }
    }
</script>

{#if active}
    <div class="flex:column" style="text-align: left; gap: 0.2em;">
        <label for="tutorToken">Token:</label>
        <input type="text" id="tutorToken" bind:this={inputRef}>
        <div class="flex:row" style="gap: 0.6em">
            <button class="filled-btn green-btn" onclick={loginClick}>Anmelden</button>
            <button class="filled-btn red-btn" onclick={() => active = false}>Abbrechen</button>
        </div>
    </div>
{:else}
    <button class="filled-btn white-btn" onclick={() => active = true}>Tutor Anmeldung</button>
{/if}


<style>
    #tutorToken {
        background-color: transparent;
        color: var(--main-fg-color);
        border: 2px solid var(--main-fg-color);
    }
</style>