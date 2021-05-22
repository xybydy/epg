<template>
  <Dialog
    v-model:visible="visible"
    closable
    header="Add New Channel"
    closeOnEscape
    modal
    dismissableMask
    @hide="onDialogHide"
    ><div class="p-field p-d-flex p-flex-column">
      <AutoComplete
        v-model="selectedCountry"
        class="p-mt-3"
        :class="[country_inv ? 'p-invalid' : '']"
        :dropdown="true"
        :suggestions="filteredCountries"
        placeholder="Country"
        @complete="getCountry($event)"
      ></AutoComplete>
      <div class="p-d-flex p-jc-end">
        <Button class="p-mr-2 p-mt-3"><i class="pi pi-times" @click="onEsc"></i></Button
        ><Button class="p-mt-3"><i class="pi pi-check" @click="onOk"></i></Button>
      </div></div
  ></Dialog>
</template>

<script setup>
  import {ref, onMounted, computed} from 'vue'
  import eventBus from '@/emitter/eventBus.js'
  import {selectedItems} from '@/store/selectedItems.js'

  import countries from '@/assets/countries'
  import {UpperFirst} from '@/utils'

  import Dialog from 'primevue/dialog'
  import Button from 'primevue/button'
  import AutoComplete from 'primevue/autocomplete'

  let visible = ref(false)
  let selectedCountry = ref('')
  let filteredCountries = ref()

  let country_inv = computed(() => {
    return selectedCountry.value === '' ? true : false
  })

  const onEsc = () => {
    visible.value = false
  }

  const onDialogHide = () => {
    selectedCountry.value = ''
  }

  const getCountry = (event) => {
    setTimeout(() => {
      filteredCountries.value =
        event.query.trim().length === 0
          ? [...countries]
          : countries.filter((c) => {
              return c.toLowerCase().startsWith(event.query.toLowerCase())
            })
    }, 250)
  }

  const onOk = () => {
    eventBus.emit('selectedCountry', selectedCountry.value)
    visible.value = false
  }

  onMounted(() => {
    eventBus.on('openMatchDialog', () => {
      visible.value = true
    })
  })
</script>
