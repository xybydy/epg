import { reactive, computed } from 'vue'

const state = reactive({
  m3uData: [],
  m3uDataLength: 0,
  selectedItems: [],
  selectedItemsLength: 0,
})

//MUTATIONS
function PUSH_DATA(data) {
  state.m3uData = data
  state.m3uDataLength = data.length
}

function PUSH_SELECTED(data) {
  state.selectedItems = data
  state.selectedItemsLength = data.length
}

// ACTIONS
async function AddM3uData(data) {
  PUSH_DATA(data)
}

async function AddSelectedItems(data) {
  PUSH_SELECTED(data)
}

// GETTERS
const GetM3uData = computed(() => state.m3uData)
const LengthM3uData = computed(() => state.m3uDataLength)
const GetSelectedItems = computed(() => state.selectedItems)
const LengthSelectedItems = computed(() => state.selectedItemsLength)

export { GetM3uData, AddM3uData, LengthSelectedItems, LengthM3uData }
