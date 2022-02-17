import { NH5, NSpace } from 'naive-ui';
import { defineComponent } from 'vue';

export default defineComponent({
  setup() {
    return () => (
      <div class="ml-4">
        <NSpace size={'small'} class="items-center">
          <img
            class="h-8 w-auto"
            alt="logo"
            src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image"
          />
          <span class="text-lg font-extrabold select-none">Echo Pro</span>
        </NSpace>
      </div>
    );
  },
});
