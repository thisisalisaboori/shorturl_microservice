using Grpc.Core;
using Grpc.Net.Client;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.



builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var channel = GrpcChannel.ForAddress("http://localhost:8080");
var client = new ShortLink.ShortLinkClient(channel);
builder.Services.AddScoped<ShortLink.ShortLinkClient>(x=> new ShortLink.ShortLinkClient(channel));

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
