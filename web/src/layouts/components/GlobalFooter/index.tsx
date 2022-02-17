import { defineComponent } from 'vue';
import { NLayoutFooter, NText } from 'naive-ui';

export default defineComponent({
  setup() {
    return () => (
      <NLayoutFooter bordered class="text-center p-4">
        <NText depth={3}>
          Copyright @ 2021-{new Date().getFullYear()} Echo Pro. All Rights
          Reserved
        </NText>
      </NLayoutFooter>
    );
  },
});
