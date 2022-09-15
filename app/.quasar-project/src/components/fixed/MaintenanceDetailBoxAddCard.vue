<template>
  <q-card class="full-width">
    <q-list dense>
      <!-- name -->
      <q-item>
        <q-item-section>
          <q-item-label>Name</q-item-label>
        </q-item-section>
        <q-item-section side>
          <q-input
            :model-value="name"
            @update:model-value="$emit('update:name', $event)"
            maxlength="45"
            dense
            hide-bottom-space
            autofocus
            :rules="[(val) => (val && val.length > 0) || 'please enter Name']"
          />
        </q-item-section>
      </q-item>
      <!-- value -->
      <q-item>
        <q-item-section>
          <q-item-label>Value</q-item-label>
        </q-item-section>
        <q-item-section side>
          <q-input
            :model-value="value"
            @update:model-value="$emit('update:value', Number($event))"
            maxlength="6"
            dense
            hide-bottom-space
            :rules="[(val) => val > -1 || 'please enter Value']"
          />
        </q-item-section>
      </q-item>
      <!-- content -->
      <q-item>
        <q-item-section>
          <q-item-label>Content</q-item-label>
        </q-item-section>
        <q-item-section side>
          <q-input
            :model-value="content"
            @update:model-value="$emit('update:content', $event)"
            type="textarea"
            maxlength="150"
            rows="2"
            dense
            counter
          />
        </q-item-section>
      </q-item>
      <!-- actions -->
      <q-item>
        <slot name="actions" />
      </q-item>
    </q-list>
  </q-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'MaintenanceDetailBoxAddCard',
  props: {
    name: {
      type: String,
      default: '',
    },
    value: {
      type: Number,
      default: 0,
    },
    content: {
      type: String,
    },
  },
  setup() {
    return {};
  },
});
</script>
