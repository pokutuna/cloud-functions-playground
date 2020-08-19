package functions;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.cloud.functions.HttpFunction;
import com.google.cloud.functions.HttpRequest;
import com.google.cloud.functions.HttpResponse;
import java.io.BufferedWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class App implements HttpFunction {
  @Override
  public void service(HttpRequest request, HttpResponse response)
      throws IOException {

    Payload p = new Payload();
    p.runtime = System.getenv("GCF_RUNTIME");
    p.key = "value";
    p.array = List.of(1,2,3);

    ObjectMapper mapper = new ObjectMapper();
    String json = mapper.writeValueAsString(p);
    System.out.println(json);

    BufferedWriter writer = response.getWriter();
    writer.write(json);
  }
}
