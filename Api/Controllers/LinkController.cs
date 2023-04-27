using Microsoft.AspNetCore.Mvc;
using Grpc.Core;
using Grpc.Net.Client;

namespace Api.Controllers;

[ApiController]
[Route("link")]
public class LinkController : ControllerBase
{
    

    private readonly ILogger<LinkController> _logger;
    private readonly ShortLink.ShortLinkClient client;

    public LinkController (ShortLink.ShortLinkClient _client,ILogger<LinkController> logger)
    {
        _logger = logger;
        client = _client;
    }

    [HttpPost()]
    public async Task<ActionResult<Response>> Post([FromBody] CreateRequest data)
    {

        var reply = await client.CreateAsync( data);
        //Console.WriteLine(reply.Shorturl);
        return reply;
    }

    [HttpGet()]
    public async Task<ActionResult<Response>> Get(string url)
    {
    

        var reply = await client.GetAsync( new GetLink(){Url =url} );
        //Console.WriteLine(reply.Shorturl);
        return reply;
    }
}
