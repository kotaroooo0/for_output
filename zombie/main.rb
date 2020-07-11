# frozen_string_literal: true

if fork
  sleep(100_000)
  p('parent process done')
else
  if fork
    sleep(20)
    p('first child process done')
  else
    if fork
      sleep(30)
      p('seocnd child process done')

    else
      sleep(40)
      p('third child process done')
    end
  end
  end

# if child_pid = fork
#   puts "parent process. pid: #{Process.pid}, child pid: #{child_pid}"

#   10.times do
#     p('parent')
#     sleep(1)
#   end

# # 子プロセスの終了を待って終了。
# #   Process.waitpid(child_pid)
# else
#   puts "child process. pid: #{Process.pid}"

#   20.times do
#     p('child')
#     sleep(1)
#   end

#   if fork
#     p('clild child')
#     sleep(3)
#   end
# end

def generate_child_process
  if fork
    sleep(10)
  else
    if fork
      sleep(20)
    else
      if fork
        sleep(30)
      else
        sleep(40)
      end
    end
  end
end
