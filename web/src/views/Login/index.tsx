import {
  FormInst,
  FormRules,
  NButton,
  NCard,
  NCheckbox,
  NForm,
  NFormItem,
  NGradientText,
  NInput,
  NSpace,
} from 'naive-ui';
import { defineComponent, reactive, ref } from 'vue';

const formRef = ref<FormInst | null>(null);

const loginRequest = reactive({
  username: '',
  password: '',
  remember_me: false,
});

const rules: FormRules = reactive({
  username: [
    { required: true, message: '用户名称必须输入', trigger: ['blur'] },
  ],
  password: [
    { required: true, message: '用户密码必须输入', trigger: ['blur'] },
  ],
});

const handleSubmit = (e: MouseEvent) => {
  if (!formRef.value) return;
  e.preventDefault();

  formRef.value.validate((errors) => {
    if (!errors) {
      // 登录
    }
  });
};

export default defineComponent({
  setup() {
    return () => (
      <div class="w-480px">
        <NCard
          // bordered={false}
          class="items-center"
          size="large"
          v-slots={{
            header: () => (
              <div class="flex justify-between">
                <NGradientText type="primary" size="28">
                  sfsdf
                </NGradientText>
              </div>
            ),
            default: () => (
              <div class="w-360px">
                <h3>登录</h3>
                <div>
                  <NForm
                    ref={formRef}
                    rules={rules}
                    size="large"
                    showLabel={false}
                  >
                    <NFormItem path="loginRequest.username" label="用户名称">
                      <NInput
                        v-model:value={loginRequest.username}
                        placeholder="请输入用户名称"
                      />
                    </NFormItem>
                    <NFormItem path="loginRequest.password" label="用户密码">
                      <NInput
                        type="password"
                        v-model:value={loginRequest.password}
                        placeholder="请输入用户密码"
                      />
                    </NFormItem>
                    <NSpace vertical={true} size={24}>
                      <div class="flex flex-row justify-between">
                        <NCheckbox v-model:checked={loginRequest.remember_me}>
                          记住我
                        </NCheckbox>
                        <NButton text={true}>忘记密码</NButton>
                      </div>
                      <NButton
                        type="primary"
                        size="large"
                        block={true}
                        round={true}
                        onClick={handleSubmit}
                      >
                        确定
                      </NButton>
                    </NSpace>
                  </NForm>
                </div>
              </div>
            ),
          }}
        />
      </div>
    );
  },
});
