import { defineComponent } from 'vue';
import { NLayoutFooter, NText } from 'naive-ui';

export default defineComponent({
  setup() {
    return () => (
      <NLayoutFooter bordered class="text-center p-4">
        <NText class="text-gray-400">
          Copyright @ 2021-{new Date().getFullYear()} Echo Pro. All Rights
          Reserved
        </NText>
      </NLayoutFooter>
    );
  },
});
