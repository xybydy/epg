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
      <InputText
        v-model="chan_name"
        :class="[chan_inv ? 'p-invalid p-mt-2' : 'p-mt-2']"
        placeholder="Channel Name"
        autofocus
      ></InputText
      ><InputText
        v-model="tvg_id"
        class="p-mt-3"
        :class="[tvg_inv ? 'p-invalid' : '']"
        placeholder="Tvg-ID"
      ></InputText
      ><AutoComplete
        v-model="selectedCountry"
        class="p-mt-3"
        :class="[country_inv ? 'p-invalid' : '']"
        :dropdown="true"
        :suggestions="filteredCountries"
        placeholder="Country"
        @complete="getCountry($event)"
      ></AutoComplete
      ><InputText v-model="logo" class="p-mt-3" placeholder="Logo URL"></InputText>
      <div class="p-d-flex p-jc-end">
        <Button class="p-mr-2 p-mt-3"><i class="pi pi-times" @click="onEsc"></i></Button
        ><Button class="p-mt-3"><i class="pi pi-check" @click="onOk"></i></Button>
      </div></div
  ></Dialog>
</template>

<script setup>
  import {ref, onMounted, computed} from 'vue'
  import eventBus from '@/emitter/eventBus.js'

  import countries from '@/assets/countries'
  import {UpperFirst} from '@/utils'

  import Dialog from 'primevue/dialog'
  import InputText from 'primevue/inputtext'
  import Button from 'primevue/button'
  import AutoComplete from 'primevue/autocomplete'

  let visible = ref(false)
  let chan_name = ref('')
  let tvg_id = ref('')
  let selectedCountry = ref('')
  let filteredCountries = ref()
  let logo = ref('')

  let chan_inv = computed(() => {
    return chan_name.value === '' ? true : false
  })
  let country_inv = computed(() => {
    return selectedCountry.value === '' ? true : false
  })
  let tvg_inv = computed(() => {
    return tvg_id.value === '' ? true : false
  })

  const onEsc = () => {
    visible.value = false
  }

  const onDialogHide = () => {
    chan_name.value = ''
    selectedCountry.value = ''
    tvg_id.value = ''
    logo.value = ''
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
    eventBus.emit('adminNewChanSend', {
      name: chan_name.value,
      tvg_id: tvg_id.value,
      country: UpperFirst(selectedCountry.value),
      logo: logo.value,
    })
    visible.value = false
  }

  onMounted(() => {
    eventBus.on('adminNewChanDialog', () => {
      visible.value = true
    })
  })
</script>
