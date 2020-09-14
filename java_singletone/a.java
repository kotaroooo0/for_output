class Main {
    public static void main(String[] args) {
        NormalSingleton n = NormalSingleton.getInstance();
        n.echo();

        EnumSingleton e = EnumSingleton.getInstance();
        e.echo();
    }
}

class NormalSingleton {
    private static final NormalSingleton INSTANCE = new NormalSingleton();

    private NormalSingleton() {
    }

    public static NormalSingleton getInstance() {
        return INSTANCE;
    }

    public void echo() {
        System.out.println("hoge");
    }
}

enum EnumSingleton {
    INSTANCE;

    private String text;

    public static EnumSingleton getInstance() {
        if (INSTANCE.text == null) {
            INSTANCE.text = "field";
        }
        return INSTANCE;
    }

    public void echo() {
        System.out.println("hoge");
    }
}
