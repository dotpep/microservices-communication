using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading.Tasks;
using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using PlatformService.AsyncDataServices;
using PlatformService.Data;
using PlatformService.Dtos;
using PlatformService.Models;
using PlatformService.SyncDataServices.Http;

namespace PlatformService.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class PlatformsController : ControllerBase
    {
        private readonly IPlatformRepo _repository;
        private readonly IMapper _mapper;
        private readonly ICommandDataClient _commandDataClient;
        private readonly IMessageBusClient _messageBusClient;

        // Injection
        public PlatformsController(
            IPlatformRepo repository,
            IMapper mapper,
            ICommandDataClient commandDataClient,
            IMessageBusClient messageBusClient
        )
        {
            _repository = repository;
            _mapper = mapper;
            _commandDataClient = commandDataClient;
            _messageBusClient = messageBusClient;
        }

        [HttpGet]
        public ActionResult<IEnumerable<PlatformReadDto>> GetPlatfroms()
        {
            Console.WriteLine("--> Getting Platforms....");

            var platformItemsList = _repository.GetAllPlatforms();

            return Ok(_mapper.Map<IEnumerable<PlatformReadDto>>(platformItemsList));
        }

        [HttpGet("{id}", Name = "GetPlatformById")]
        public ActionResult<PlatformReadDto> GetPlatformById(int id)
        {
            var platformSpecificOneItem = _repository.GetPlatformById(id);

            if (platformSpecificOneItem != null)
            {
                return Ok(_mapper.Map<PlatformReadDto>(platformSpecificOneItem));
            }

            return NotFound();
        }

        [HttpPost]
        public async Task<ActionResult<PlatformReadDto>> CreatePlatform(PlatformCreateDto platformCreateDto)
        {
            var platformModel = _mapper.Map<Platform>(platformCreateDto);
            _repository.CreatePlatform(platformModel);
            _repository.SaveChanges();

            // wrapp Dto with mapper to Model
            var platformReadDto = _mapper.Map<PlatformReadDto>(platformModel);
            

            // Http Client - Send Sync Message
            try
            {
                await _commandDataClient.SendPlatformToCommand(platformReadDto);
            }
            catch (Exception ex)
            {
                Console.WriteLine($"--> Could not send Syncronously: {ex.Message}");
            }

            // Message Broker Client - Send Async Message
            try
            {
                var platformPublishedDto = _mapper.Map<PlatformPublishedDto>(platformReadDto);
                platformPublishedDto.Event = "Platform_Published";
                _messageBusClient.PublishNewPlatform(platformPublishedDto);
            }
            catch (Exception ex)
            {
                Console.WriteLine($"--> Could not send Asyncronously: {ex.Message}");
            }

            return CreatedAtRoute(
                nameof(GetPlatformById),
                new { Id = platformReadDto.Id },
                platformReadDto
            );
        }
    }
}