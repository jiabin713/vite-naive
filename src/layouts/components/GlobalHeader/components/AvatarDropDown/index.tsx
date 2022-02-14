import { LogOutOutline, PersonCircleOutline } from '@vicons/ionicons5';
import { NButton, NDropdown, NIcon } from 'naive-ui';
import { defineComponent, h } from 'vue';

const renderIcon = (icon: any) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const options = [
  {
    label: '用户资料',
    key: 'profile',
    icon: renderIcon(PersonCircleOutline),
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogOutOutline),
  },
];

export default defineComponent({
  setup() {
    return () => (
      <NDropdown options={options}>
        <NButton
          text
          ghost
          secondary
          v-slots={{
            icon: () => <PersonCircleOutline />,
            default: () => <span>个人中心</span>,
          }}
        />
      </NDropdown>
    );
  },
});
