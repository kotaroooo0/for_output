import java.util.List;

public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello World!");
        String hoge = "hoge";
        System.out.println(hoge.length());

        General<String> general = new General<>();
        general.put("object");
        String fuga = general.get();
        System.out.println(fuga);

        List<String> ll1 = List.of("1","2","3");
        List<String> ll2 = List.of("1","2","3");
        System.out.println(numElementInCommon(ll1, ll2));

    }

    static int numElementInCommon(List<?> l1, List<?> l2) {
        int result = 0;
        for (Object o : l1) {
            if (l2.contains(o)) {
                result++;
            }
        }
        return result;
    }
}


class General<T> {
    T object;

    void put(T object) {
        this.object = object;
    }

    T get() {
        return object;
    }

}
