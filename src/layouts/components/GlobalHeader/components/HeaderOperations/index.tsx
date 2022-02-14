import { NButton, NIcon, NSpace, NTooltip } from 'naive-ui';
import { defineComponent } from 'vue';
import {
  InvertModeOutline,
  MailOutline,
  SearchOutline,
} from '@vicons/ionicons5';
import AvatarDropDown from '../AvatarDropDown';

export default defineComponent({
  setup() {
    return () => (
      <div class="mr-4">
        <NSpace wrap={false} size={'large'} class="justify-end items-center">
          <NTooltip
            placement="bottom"
            trigger="hover"
            v-slots={{
              trigger: () => (
                <NButton
                  text
                  circle
                  ghost
                  secondary
                  v-slots={{
                    icon: () => (
                      <NIcon>
                        <SearchOutline />
                      </NIcon>
                    ),
                  }}
                />
              ),
              default: () => <span>搜索</span>,
            }}
          />
          <NTooltip
            placement="bottom"
            trigger="hover"
            v-slots={{
              trigger: () => (
                <NButton
                  text
                  circle
                  ghost
                  secondary
                  v-slots={{
                    icon: () => (
                      <NIcon>
                        <MailOutline />
                      </NIcon>
                    ),
                  }}
                />
              ),
              default: () => <span>搜索</span>,
            }}
          />
          <NTooltip
            placement="bottom"
            trigger="hover"
            v-slots={{
              trigger: () => (
                <NButton
                  text
                  circle
                  ghost
                  secondary
                  v-slots={{
                    icon: () => (
                      <NIcon>
                        <InvertModeOutline />
                      </NIcon>
                    ),
                  }}
                />
              ),
              default: () => <span>深色模式</span>,
            }}
          />
          <AvatarDropDown />
        </NSpace>
      </div>
    );
  },
});
