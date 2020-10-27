<template>
  <DataTable
    class="p-datatable-sm p-datatable-striped editable-cells-table"
    editMode="cell"
    :value="m3u"
    :paginator="true"
    :rows="10"
  >
    <Column field="name" header="NAME">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" />
      </template>
    </Column>
    <Column field="group.title" header="GROUP">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
        /> </template
    ></Column>
    <Column field="tvg.id" header="TVG-ID">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
        /> </template
    ></Column>
    <Column field="tvg.name" header="TVG-NAME">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
        /> </template
    ></Column>
    <Column field="tvg.logo" header="TVG-LOGO">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
        /> </template
    ></Column>
    <Column field="url" header="URL">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
        /> </template
    ></Column>
  </DataTable>
</template>

<script>
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
//import ColumnGroup from 'primevue/columngroup'

export default {
  props: {
    m3u: Object
  },
  setup() {
    return {}
  },
  components: {
    DataTable,
    Column
  }
}
</script>

<style lang="scss" scoped></style>
