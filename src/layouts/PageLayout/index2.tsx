import { NLayout } from 'naive-ui';
import { defineComponent } from 'vue';
import GlobalFooter from '../components/GlobalFooter';
import GlobalHeader from '../components/GlobalHeader';
import GlobalSider from '../components/GlobalSider';
import PageContainer from '../components/PageContainer';

export default defineComponent({
  setup() {
    return () => (
      <div class="flex flex-col h-screen">
        <div class="h-16 bg-gray-400 flex-shrink-0">header</div>
        <div class="flex overflow-y-auto h-full flex-row flex-nowrap">
          <div class="w-16rem bg-red-100 ">side</div>
          <div class="flex-1 overflow-y-auto flex flex-col">
            <div class="bg-yellow-200 flex-grow">content</div>
            <div class="h-16 bg-gray-800 flex-none">footer</div>
          </div>
        </div>
      </div>
    );
  },
});
