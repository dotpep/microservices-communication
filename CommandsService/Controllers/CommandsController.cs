using System;
using System.Collections.Generic;
using AutoMapper;
using CommandsService.Data;
using CommandsService.Dtos;
using CommandsService.Models;
using Microsoft.AspNetCore.Mvc;

namespace CommandsService.Controllers
{
    [Route("api/c/platforms/{platformId}/[controller]")]
    [ApiController]
    public class CommandsController : ControllerBase
    {
        private readonly ICommandRepo _repository;
        private readonly IMapper _mapper;

        // Injection
        public CommandsController(ICommandRepo repository, IMapper mapper)
        {
            _repository = repository;
            _mapper = mapper;
        }

        [HttpGet]
        public ActionResult<IEnumerable<CommandReadDto>> GetListOfCommandsByPlatformId(int platformId)
        {
            Console.WriteLine("--> Getting Commands for specific Platform by Id");
            Console.WriteLine($"--> Hit GetListOfCommandsByPlatformId: platformId:{platformId}");
            
            if (!_repository.PlatformExists(platformId))
            {
                return NotFound();
            }

            var commandItems = _repository.GetAllCommandsForPlatform(platformId);

            if (commandItems == null)
            {
                return NotFound();
            }

            return Ok(_mapper.Map<IEnumerable<CommandReadDto>>(commandItems));
        }

        [HttpGet("{commandId}", Name = "GetCommandForPlatform")]
        public ActionResult<CommandReadDto> GetCommandForPlatform(int platformId, int commandId)
        {
            Console.WriteLine("--> Getting Specific Command for Specific Platform by Id's");
            Console.WriteLine($"--> Hit GetCommandForPlatform: platformId:{platformId} / commandId:{commandId}");

            if (!_repository.PlatformExists(platformId))
            {
                return NotFound();
            }

            var commandForPlatformItem = _repository.GetCommand(platformId, commandId);

            if (commandForPlatformItem == null)
            {
                return NotFound();
            }

            return Ok(_mapper.Map<CommandReadDto>(commandForPlatformItem));
        }

        [HttpPost]
        public ActionResult<CommandReadDto> CreateCommandForPlatform(int platformId, CommandCreateDto commandCreateDto)
        {
            Console.WriteLine("--> Creating Command for Specific Platform by platformId");
            Console.WriteLine($"--> Hit CreateCommandForPlatform: platformId:{platformId}");

            if (!_repository.PlatformExists(platformId))
            {
                return NotFound();
            }

            var commandModel = _mapper.Map<Command>(commandCreateDto);
            _repository.CreateCommand(platformId, commandModel);
            _repository.SaveChanges();

            // wrapp Dto with mapper to Model
            var commandReadDto = _mapper.Map<CommandReadDto>(commandModel);

            return CreatedAtRoute(
                nameof(GetCommandForPlatform),
                new { platformId = platformId, commandId = commandReadDto.Id },
                commandReadDto
            );
        }
    }
}